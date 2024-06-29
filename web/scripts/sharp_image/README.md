# 影像壓縮 SHARP

檔案放 `src` 資料夾

run `node index.js`

輸出在 `compressed_images` 資料夾

調整編碼格式需要改code

https://github.com/lovell/sharp

## HEIC轉欓支援

sharp預設不支援heic

但可以透過全域安裝上游編碼套件來支援
```
brew install libheif
brew install libvips
pnpm remove sharp && pnpm add sharp
```
