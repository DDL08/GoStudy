#导出IP出现次数

from scapy.all import rdpcap, IP

input_file = r"C:\Users\LEGION\Desktop\10.10.10.12\202505251119\out.pcap"
packets = rdpcap(input_file)

ip_count = {}

for pkt in packets:
    if IP in pkt:
        src_ip = pkt[IP].src
        dst_ip = pkt[IP].dst
        ip_count[src_ip] = ip_count.get(src_ip, 0) + 1
        ip_count[dst_ip] = ip_count.get(dst_ip, 0) + 1

# 按次数排序输出
for ip, count in sorted(ip_count.items(), key=lambda x: x[1], reverse=True):
    print(f"{ip}: {count}")
