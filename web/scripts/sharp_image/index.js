const fs = require("fs");
const path = require("path");
const sharp = require("sharp");

// 压缩图片的目标文件夹
const outputFolder = path.join(process.env.HOME, "downloads/sharp");
// const outputFolder = "./compressed_images";

// 确保输出文件夹存在
if (!fs.existsSync(outputFolder)) {
  fs.mkdirSync(outputFolder);
}

// 获取当前文件夹中的所有文件
const inputFolder = "./src";
const files = fs.readdirSync(inputFolder);

console.log({ sourceImages: files });

// 过滤出图片文件
const imageFiles = files.filter((file) => {
  const ext = path.extname(file).toLowerCase();
  return [".jpg", ".jpeg", ".png", ".gif", ".webp", ".tiff", ".heic"].includes(
    ext
  );
});

// 遍历图片文件并压缩
imageFiles.forEach((file) => {
  const inputPath = path.join(__dirname, inputFolder, file);
  const outputPath = path.join(
    outputFolder,

    // file
    path.parse(file).name + "." + "png"
    // path.parse(file).name + "." + "jpg"
    // path.parse(file).name + "." + "webp"
  );

  const outputType = outputPath.split(".").at(-1);
  const sharpTypeName = {
    jpg: "jpeg",
    png: "png",
    webp: "webp",
  }[outputType];

  sharp(inputPath)
    ?.[sharpTypeName]()
    // .resize({ width: 1280 }) // 设置压缩后的宽度，可根据需要调整
    // .resize({ width: 400, height: 400 })
    .toFile(outputPath, (err, info) => {
      if (err) {
        console.error(`Failed to compress ${file}: ${err}`);
      } else {
        console.log(`Compressed ${file} successfully`);
      }
    });
});
