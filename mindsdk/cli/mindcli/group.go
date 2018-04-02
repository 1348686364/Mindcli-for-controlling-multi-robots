package mindcli

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func (mindcli *MindCli) AllocatePort(robotIP string) {
	//run skill in this robot on a avaliable port
	robotsport := mindcli.portConfig.Robotsport
	var robotport Robotport
	robotport.RobotIp = robotIP
	flag1 := 0
	flag2 := 0
	for flag1 != 1 || flag2 != 1 {
		if !IsUDPPortAvailable(mindcli.config.ServeMPKPort) || !IsTCPPortAvailable(mindcli.config.ServeMPKPort) {
			//fmt.Printf("ServeMPKPort %d already in use\n we allocate it as port", mindcli.config.ServeMPKPort)
			mindcli.config.ServeMPKPort += 1
		} else {
			robotport.ServeMPKPort = mindcli.config.ServeMPKPort
			fmt.Printf("ServeMPKPort of %s is: %d \n", robotIP, robotport.ServeMPKPort)
			flag1 = 1
		}
		if !IsUDPPortAvailable(mindcli.config.ServeRemotePort) || !IsTCPPortAvailable(mindcli.config.ServeRemotePort) {
			//fmt.Printf("ServeRemotePort %d already in use\n we allocate it as port", mindcli.config.ServeRemotePort)
			mindcli.config.ServeRemotePort += 1
		} else {
			robotport.ServeRemotePort = mindcli.config.ServeRemotePort
			fmt.Printf("ServeRemotePort of %s is: %d \n", robotIP, robotport.ServeRemotePort)
			flag2 = 1
		}
	}
	for i := 0; i < len(mindcli.portConfig.Robotsport); i++ {
		if mindcli.portConfig.Robotsport[i].RobotIp == robotport.RobotIp {
			mindcli.portConfig.Robotsport[i] = robotport
			flag1 = 0
		}
	}
	if flag1 == 1 {
		mindcli.portConfig.Robotsport = append(robotsport, robotport)
	}
	mindcli.portConfig.Write()
}

//group manger:AddRobotToGroup
func (mindcli *MindCli) AddRobotToGroup(robotname string, groupname string) error {
	var robname []string
	for i := 0; i < len(mindcli.userConfig.Groups); i++ {
		if mindcli.userConfig.Groups[i].Name == groupname {
			robname = mindcli.userConfig.Groups[i].Robotsname
			mindcli.userConfig.Groups[i].Robotsname = append(robname, robotname)
			mindcli.userConfig.Write()
			return nil
		}
	}
	return errors.New("Could not find group with that name, plase create a new group by [addg]!")
}

//group manger:AddGroup
func (mindcli *MindCli) AddGroup(groupname string) error {
	//make sure the group name is unique
	for _, group := range mindcli.userConfig.Groups {
		if group.Name == groupname {
			return errors.New("This group is already created, the name of groups shuold be unique!")
		}
	}
	//add group
	var group_new Group
	var groups_backup []Group
	group_new.Name = groupname
	groups_backup = mindcli.userConfig.Groups
	mindcli.userConfig.Groups = append(groups_backup, group_new)
	mindcli.userConfig.Write()
	return nil
}

//group manger:DeleteRobotFromGroup
func (mindcli *MindCli) DeleteRobotFromGroup(robotname string, groupname string) error {
	var robotsname_fi []string //this var will be the same as the robotsname in target group
	var flag int
	flag = 0
	// flag = 0: cannot Find group;
	// flag = 1: cannot find robot ;
	// flag = 2: find the target
	for i := 0; i < len(mindcli.userConfig.Groups); i++ {
		//find the target group
		if mindcli.userConfig.Groups[i].Name == groupname {
			flag = 1
			for j := 0; j < len(mindcli.userConfig.Groups[i].Robotsname); j++ {
				//copy all robotname except the target robotname
				if mindcli.userConfig.Groups[i].Robotsname[j] == robotname {
					flag = 2
				} else {
					robotsname_fin := append(robotsname_fi, mindcli.userConfig.Groups[i].Robotsname[j])
					robotsname_fi = robotsname_fin
				}
			}
			// update the robotsname in mindcli.userConfig.Groups[i]
			mindcli.userConfig.Groups[i].Robotsname = robotsname_fi
		}
	}
	//check the flag and decide whether the final of []robot should be saved
	if flag == 0 {
		return errors.New("Cannot find the group please check the group name!")
	} else if flag == 1 {
		return errors.New("Cannot find robot in your target group, please check the robot and group name!")
	} else if flag == 2 {
		// only when flage is 2 we save the change
		mindcli.userConfig.Write()
		return nil
	} else {
		return errors.New("The value of flag is illegal")
	}
}

//group manger:DeleteGroup
func (mindcli *MindCli) DeleteGroup(groupname string) error {
	//
	var group_fi []Group //this var will be the same as the robotsname in target group
	var flag int
	flag = 0
	// flag = 0: cannot Find group;
	// flag = 1: find the target group;
	for i := 0; i < len(mindcli.userConfig.Groups); i++ {
		if mindcli.userConfig.Groups[i].Name == groupname {
			flag = 1
		} else {
			//copy all group except the target robotname
			group_fin := append(group_fi, mindcli.userConfig.Groups[i])
			group_fi = group_fin
		}
	}
	// update the robotsname in mindcli.userConfig.Groups[i]
	mindcli.userConfig.Groups = group_fi
	//check the flag and decide whether the final []group should be saved
	if flag == 0 {
		return errors.New("Cannot find the group, please check the group name!")
	} else if flag == 1 {
		// only when flage is 1 we save the change
		mindcli.userConfig.Write()
		return nil
	} else {
		return errors.New("The value of flag is illegal")
	}
}

//group manger:ListRobotInGroup
func (mindcli *MindCli) ListRobotInGroup(groupname string) error {
	//make sure the group is exist
	for _, group := range mindcli.userConfig.Groups {
		if group.Name == groupname {
			//get all of the robot name from group
			//fmt.Printf("	group		robot\n ")
			fmt.Printf("			%s 			\n", groupname)
			for _, robot := range group.Robotsname {
				fmt.Printf("	groupname:%s		robotname：%s\n ", groupname, robot)
			}
			return nil
		}
	}
	return errors.New("Cannot find the group, please check the group name!")
}

//group manger:ListGroup
func (mindcli *MindCli) ListGroup() error {
	// find if there is any groups
	if len(mindcli.userConfig.Groups) > 0 {
		fmt.Print("The groups you have created are:\n")
		for _, group := range mindcli.userConfig.Groups {
			err := mindcli.ListRobotInGroup(group.Name)
			if err != nil {
				fmt.Println(err)
				os.Exit(-1)
			}
		}
		fmt.Print("List group finshed, this is all.\n")
		return nil
	}
	return errors.New("Cannot find any groups.")
}

func (mindcli *MindCli) execrun(robotIP string, noInstall bool) *exec.Cmd {
	cmd := exec.Command("mind", "run", "--ip", robotIP)
	if noInstall {
		cmd = exec.Command("mind", "run", "--ip", robotIP, "-n")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	time.Sleep(time.Second)
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	// cmd.Wait()
	return cmd
}

func (mindcli *MindCli) execscan() {
	cmd := exec.Command("mind", "scan")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func (mindcli *MindCli) GroupRun_n(groupname string) error {
	err := mindcli.GroupRun_base(groupname, true)
	return err
}

func (mindcli *MindCli) GroupRun(groupname string) error {
	err := mindcli.GroupRun_base(groupname, false)
	return err
}

func (mindcli *MindCli) GroupRun_base(groupname string, noInstall bool) error {
	exec.Command("cd")
	exec.Command("rm", ".mindport.json")
	var flag int
	flag = 0
	var robot Robot
	var cmds []*exec.Cmd
	mindcli.execscan()

	// make sure the group is exist
	for i := 0; i < len(mindcli.userConfig.Groups); i++ {

		if mindcli.userConfig.Groups[i].Name == groupname {
			// run skill in each robot
			//get all of the robots' name from group

			for _, robotname := range mindcli.userConfig.Groups[i].Robotsname {
				// find robot from userConfig
				for _, robot_usr := range mindcli.userConfig.Robots {
					if robot_usr.Name == robotname {
						robot = robot_usr
						flag = 2
						break
					}
				}
				//onece there is no such a robotnaem don't do the following steps
				if flag == 2 {
					fmt.Printf("Run robot %s \n", robot.Name)
					cmd := mindcli.execrun(robot.IP, noInstall)
					cmds = append(cmds, cmd)
				} else {
					fmt.Printf("Cannot find robot %s in user config.\n", robotname)
				}
				flag = 1
			}
		}
	}
	for _, cmd := range cmds {
		cmd.Wait()
	}
	if flag == 1 {
		return nil
	} else {
		return errors.New("Cannot find the group or there is no robots in group, please check the group name by [listr] <groupname>!")
	}
}
