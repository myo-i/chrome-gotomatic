from selenium import webdriver
import time
 
driver = webdriver.Chrome("C:\\Users\\PC_User\\MyProject\\chromedriver_win32\\chromedriver.exe")
 
#空のタブを開く
driver.execute_script("window.open();")
 
# URLを指定して、新しいタブを開く
driver.execute_script("window.open('https://finance.yahoo.co.jp/');")

time.sleep(10)