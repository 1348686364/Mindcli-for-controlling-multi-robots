# -*- coding: utf-8 -*-
import easygui as g
import sys
import json
import os.path
from selenium import webdriver
import sys
import time

reload(sys)
sys.setdefaultencoding('utf8')

MINDPORT_FILE_NAME = '/home/zxt/.mindport.json'
MIND_FILE_NAME = '/home/zxt/.mind.json'

def GetJson(filename):
    if not os.path.isfile(filename):
        print('{} does not exit'.format(filename))
        sys.exit()
    load_dict = json.load(open(filename,'r'))
    return load_dict

# return dictionary array of robots information
def GetRobots():
    return GetJson(MIND_FILE_NAME).get('Robots')

#return string array of robots name
def GetRobotsName():
    robotname = []
    for i in GetRobots():
        robotname.append(i.get('Name'))
    return robotname
def GetRobotIp(name):
    for i in GetRobots():
        if i.get('Name') == name :
            robotip = i.get('IP')
    return robotip

# return dictionary array of ports information
def GetPortsInfo():
    return GetJson(MINDPORT_FILE_NAME).get('Robotsport')
def GetServeRemotePort(name):
    for i in GetPortsInfo():
        if i.get('RobotIp') == GetRobotIp(name) :
            ServeRemotePort = i.get('ServeRemotePort')
    return ServeRemotePort


# use shell to reach these function

def init():
    browser = webdriver.Chrome()
    browser.set_window_size(0,0)
    port_dicts = GetPortsInfo()
    urls =[]
    for port_dict in port_dicts:
        url = 'localhost:'+str(port_dict.get("ServeRemotePort"))
        urls.append(url)
    OpenNewSessions(browser,urls)
    # time.sleep(5)
    return browser,urls

def OpenNewSessions(browser,urls):
    for url in urls:
        browser.execute_script('window.open("http:' + url + '");')
        browser.implicitly_wait(30)

def ClickAll(button_id,browser):
    # since the first handle is not a url we create, we overlap it
    for i in range(len(browser.window_handles) - 1):
            browser.switch_to_window(browser.window_handles[i+1])
            # if bool(browser.find_element_by_id(button_id)):
            browser.find_element_by_id(button_id).click()

def ClickSignle(Robotname,button_id,browser,urls):
    url = 'localhost:'+ str(GetServeRemotePort(Robotname))
    browser.switch_to_window(browser.window_handles[urls.index(url)+1])
    # if bool(browser.find_element_by_id(button_id)):
    browser.find_element_by_id(button_id).click()



def exit(browser):
        browser.quit()

# design the gui for testing
def gui():
    browser,urls = init()
    #welcom page
    g.msgbox("welcom to the group HEXAController!")
    #information in Index page
    IndexMsg = "Chose your control command:"
    IndexTitle = "HEXAController"
    IndexItem = ["start all","stop all","single control","exit"]

    # information in chosing robot page
    SingleControlChoiceMsg = "Chose the robot name you want control:"
    SingleControlChoiceTitle = "RobotChose"
    SingleControlChoicesItem = GetRobotsName()

    #information in controlling single robot
    SingleControlMsg = "Chose your command to the controlled robot:"
    SingleControlTitle = "SingleController_"
    SingleControlItem = ["start","stop","exit"]

    # gui
    while 1:
        # index page
        IndexChoice = g.buttonbox(IndexMsg, IndexTitle, IndexItem )
        if IndexChoice == "start all":
            ClickAll("start",browser)
        elif IndexChoice == "stop all":
            ClickAll("stop",browser)
        elif IndexChoice == "single control":
            # new control page
            Robotname = g.choicebox(SingleControlMsg, SingleControlTitle, SingleControlChoicesItem )
            IFGoAhead = 0
            if bool(Robotname):
                IFGoAhead = 1
            while IFGoAhead:
                SignleChioce = g.buttonbox(SingleControlMsg, SingleControlTitle + Robotname, SingleControlItem)
                if SignleChioce == "start":
                    ClickSignle(Robotname,"start",browser,urls)
                elif SignleChioce == "stop":
                    ClickSignle(Robotname,"stop",browser,urls)
                elif SignleChioce == "exit":
                    break
        elif IndexChoice == "exit":
            break

gui()
