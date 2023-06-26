#!/usr/bin/env python3
# -*- coding: utf-8 -*-
from PIL import Image
from PIL.PngImagePlugin import PngInfo

import os
# assign directory
directory = './groupImgs'
 
# iterate over files in
# that directory
for filename in os.listdir(directory):
    f = os.path.join(directory, filename)
    # checking if it is a file
    if os.path.isfile(f):
        print(f)
        img = Image.open(f)

        new_img = img.transpose(Image.FLIP_LEFT_RIGHT)

        new_img.save(f, pnginfo=metadata)


import os
os.chdir('./groupImgs') 

bg = Image.new('RGB',(1200, 800), '#000000') # 產生一張 1200x800 的全黑圖片
for i in range(1,9):
    img = Image.open(f'd{i}.jpg')  # 開啟圖片
    img = img.resize((300, 400))   # 縮小尺寸為 300x400
    x = (i-1)%4                    # 根據開啟的順序，決定 x 座標
    y = (i-1)//4                   # 根據開啟的順序，決定 y 座標 ( // 為快速取整數 )
    bg.paste(img,(x*300, y*400))   # 貼上圖片

# 224x336

bg.save('./groupImgs/output/oxxostudio.jpg')
