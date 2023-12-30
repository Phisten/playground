import puppeteer from "puppeteer-extra";
import { executablePath } from "puppeteer";
import StealthPlugin from "puppeteer-extra-plugin-stealth";

export const puppeteer_demo = () => {
  puppeteer
    .use(StealthPlugin())
    .launch({
      headless: true,
      executablePath: executablePath(),
    })
    .then(async (browser) => {
      const page = await browser.newPage();
      await page.goto("https://bot.sannysoft.com");
      await page.waitForTimeout(5000);
      await page.screenshot({ path: "stealth.png", fullPage: true });
      await browser.close();
    });
};
