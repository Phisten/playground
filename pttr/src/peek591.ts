import puppeteer from "puppeteer-extra";
import { executablePath } from "puppeteer";
import StealthPlugin from "puppeteer-extra-plugin-stealth";
import dayjs from "dayjs";

const uri = `https://rent.591.com.tw/?multiPrice=5000_10000&section=3,2&searchtype=1&multiNotice=all_sex,boy&showMore=1&order=posttime&orderType=desc&other=newPost&option=broadband&firstRow=0&totalRows=44`;

export const puppeteer_peek591 = () => {
  puppeteer
    .use(StealthPlugin())
    .launch({
      headless: true,
      executablePath: executablePath(),
    })
    .then(async (browser) => {
      const page = await browser.newPage();
      await page.goto(uri);
      await page.waitForTimeout(5000);
      await page.screenshot({
        // type: "jpeg",
        path: `591-${dayjs().format("YYYY-MM-DD HH:mm")}.jpg`,
        fullPage: true,
      });

      console.log({ page });
      const itemsContent = "switch-list-content"; //物件清單
      page.$(itemsContent).then((items) => {
        console.log({ items });
      });

      await browser.close();
    });
};

puppeteer_peek591();
