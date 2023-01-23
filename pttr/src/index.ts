import puppeteer from "puppeteer-extra";
import StealthPlugin from "puppeteer-extra-plugin-stealth";

// require executablePath from puppeteer
const {executablePath} = require('puppeteer')

puppeteer
  .use(StealthPlugin())
  .launch({ headless: true,

    // add this
    executablePath: executablePath(), })
  .then(async (browser) => {
    const page = await browser.newPage();
    await page.goto("https://bot.sannysoft.com");
    await page.waitForTimeout(5000);
    await page.screenshot({ path: "stealth.png", fullPage: true });
    await browser.close();
  });
