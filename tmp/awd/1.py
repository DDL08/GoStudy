import requests
from concurrent.futures import ThreadPoolExecutor
import threading

# 创建一个线程锁
lock = threading.Lock()

# 打开文件
f = open(r'F:\addl\awd\hosts.txt', "w")

def get_ip(url):
    try:
        resp = requests.get(url, timeout=1)  # 设置超时时间
        status = resp.status_code
        if status == 200:
            with lock:  # 使用锁确保线程安全
                f.write(url + "\n")
                print(url)
    except Exception as e:
        print(f"Error accessing {url}: {e}")

url_list = []
for i in range(1, 255):
    url_list.append(f"http://192.168.1.{i}.pvp6109.bugku.cn")  # 修正URL格式

with ThreadPoolExecutor(max_workers=50) as executor:  # 减少线程池大小
    executor.map(get_ip, url_list)

# 关闭文件
f.close()