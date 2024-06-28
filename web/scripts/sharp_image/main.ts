const fs = require('fs');
const path = require('path');
const sharp = require('sharp');

// 压缩图片的目标文件夹
const outputFolder = './compressed_images';

// 确保输出文件夹存在
if (!fs.existsSync(outputFolder)) {
  fs.mkdirSync(outputFolder);
}

// 获取当前文件夹中的所有文件
const files = fs.readdirSync('./');

// 过滤出图片文件
const imageFiles = files.filter((file) => {
  const ext = path.extname(file).toLowerCase();
  return ['.jpg', '.jpeg', '.png', '.gif', '.webp', '.tiff'].includes(ext);
});

// 遍历图片文件并压缩
imageFiles.forEach((file) => {
  const inputPath = path.join(__dirname, file);
  const outputPath = path.join(__dirname, outputFolder, file);
  sharp(inputPath)
    // .resize({ width: 1020, height: 1020 }) // 设置压缩后的宽度，可根据需要调整
    .toFile(outputPath, (err, info) => {
      if (err) {
        console.error(`Failed to compress ${file}: ${err}`);
      } else {
        console.log(`Compressed ${file} successfully`);
      }
    });
});
