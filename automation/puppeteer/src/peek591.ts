import puppeteer from "puppeteer-extra";
import { ElementHandle, executablePath } from "puppeteer";
import StealthPlugin from "puppeteer-extra-plugin-stealth";
import dayjs from "dayjs";
import axios from "axios";

const uri = `https://rent.591.com.tw/?multiPrice=5000_10000&section=3,2&searchtype=1&multiNotice=all_sex,boy&showMore=1&order=posttime&orderType=desc&other=newPost&option=broadband&firstRow=0&totalRows=44`;

export const puppeteer_peek591 = () => {
  puppeteer
    .use(StealthPlugin())
    .launch({
      headless: true,
      executablePath: executablePath(),
    })
    .then(async (browser) => {
      console.log("init 591");
      const page = await browser.newPage();

      console.log("waiting page loaded");
      await page.goto(uri, { waitUntil: "networkidle0" });
      await page.waitForTimeout(4000);

      const fileNameWithoutExt = `591-${dayjs().format("YYYY-MM-DD HH:mm")}`;
      console.log(`screenshot: ${fileNameWithoutExt}`);
      await page.screenshot({
        path: `${fileNameWithoutExt}.jpg`,
        fullPage: true,
      });

      // console.log("try bodyHTML");
      // let bodyHTML = await page.evaluate(() => document.body.innerHTML);
      // console.log({ bodyHTML });

      const itemsContent = ".item-title"; //物件名 list
      console.log(`try ${itemsContent}`);
      await page
        .waitForSelector(itemsContent)
        .then(() => {
          console.log("find itemsContent");
        })
        .catch((error) => {
          console.log({ error });
        });

      await page.waitForTimeout(3000);

      const houseList = await page.evaluate(() => {
        const el = document.querySelectorAll(".item-title");

        return el;
      });
      console.log({ "find data: ": houseList });

      console.log("index 0 ", houseList[0]);

      // for (let index = 0; index < houseList.length; index++) {
      //   console.log(houseList[index]);
      // }

      // console.log({ data });

      await browser.close();
    });
};

//負責抓類別的函數，這裡是async函數必須要回傳Promise
// async function getHouse(content: string): Promise<Array<Object>> {
//   let data: Array<Object> = [];
//   return data;
// }

function saveToFile(data: Array<Object>, fileNameWithoutExt: string) {
  const fs = require("fs");
  const content = JSON.stringify(data); //轉換成json格式
  fs.writeFile(
    `${fileNameWithoutExt}.json`,
    content,
    "utf8",
    function (err: Error) {
      if (err) {
        return console.log(err);
      }
      console.log("The file was saved!");
    }
  );
}

const asd = () => {
  axios
    .get(
      "https://rent.591.com.tw/home/search/rsList?is_format_data=1&is_new_list=1&type=1&multiPrice=5000_10000&section=3,2&searchtype=1&multiNotice=all_sex,boy&showMore=1&order=posttime&orderType=desc&other=newPost&option=broadband&firstRow=0&totalRows=44"
    )
    .then((v) => {
      console.log({ v });
    })
    .catch((err) => {
      console.log({ err });
    });
};

// puppeteer_peek591();
asd();
