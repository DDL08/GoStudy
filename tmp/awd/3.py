import requests
import concurrent.futures
from urllib.parse import urlparse
import socket
import time

def check_server(ip):
    url = f"http://192-168-1-{ip}.pvp6115.bugku.cn"  # 构造URL
    try:
        # 设置超时时间，避免等待太久
        response = requests.get(url, timeout=3)
        # 如果响应状态码是200，或者服务器有响应（即使返回错误）
        if response.status_code < 500:
            return ip, True
    except requests.exceptions.RequestException:
        # 连接失败，服务器可能不存在
        pass
    
    return ip, False

def scan_network(start_ip=1, end_ip=255, max_threads=50):
    active_servers = []
    
    print(f"开始扫描192.168.1.{start_ip}到192.168.1.{end_ip}范围内的服务器...")
    
    with concurrent.futures.ThreadPoolExecutor(max_workers=max_threads) as executor:
        # 创建IP范围列表
        ip_range = range(start_ip, end_ip + 1)
        
        # 使用线程池并发检查每个IP
        future_to_ip = {executor.submit(check_server, ip): ip for ip in ip_range}
        
        for future in concurrent.futures.as_completed(future_to_ip):
            ip = future_to_ip[future]
            try:
                ip_num, is_active = future.result()
                if is_active:
                    active_ip = f"192.168.1.{ip_num}"
                    print(f"发现活跃服务器: {active_ip}")
                    active_servers.append(active_ip)
            except Exception as e:
                print(f"检查IP 192.168.1.{ip}时出错: {e}")
    
    print("\n扫描完成！发现以下活跃服务器:")
    for server in active_servers:
        print(f"http://192-168-1-{server.split('.')[-1]}.pvp6115.bugku.cn (IP: {server})")

if __name__ == "__main__":
    # 设置扫描范围，默认是1-255
    start = 1
    end = 255
    
    # 开始扫描
    scan_network(start, end)