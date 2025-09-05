#文件名：zhuan.py
import binascii
filename = r"C:\Users\LEGION\Desktop\b1.exe"
with open(filename, 'rb') as f:
    content = f.read()
    print(binascii.hexlify(content))