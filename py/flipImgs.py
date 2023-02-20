#!/usr/bin/env python3
# -*- coding: utf-8 -*-
from PIL import Image
from PIL.PngImagePlugin import PngInfo

import os
# assign directory
directory = './flipImgs_FLIP_LEFT_RIGHT'
 
# iterate over files in
# that directory
for filename in os.listdir(directory):
    f = os.path.join(directory, filename)
    # checking if it is a file
    if os.path.isfile(f):
        print(f)
        img = Image.open(f)
        metadata = PngInfo()

        new_img = img.transpose(Image.FLIP_LEFT_RIGHT)

        metadata.add_text('parameters', img.info.get('parameters'))

        new_img.save(f, pnginfo=metadata)
