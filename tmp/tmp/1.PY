import pyautogui
import time

print("请在5秒内点击到输入框...")
time.sleep(5)
text_to_type = '''123'''
pyautogui.typewrite(text_to_type, interval=0.1)

import pyautogui
pyautogui.typewirte(text,interval=0.1)