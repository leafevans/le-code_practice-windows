const puppeteer = require('puppeteer');

(async () => {
    console.log('🚀 正在启动浏览器...');

    const browser = await puppeteer.launch({
        headless: false,
        defaultViewport: null,
        args: ['--start-maximized']
    });

    const page = await browser.newPage();

    console.log('🔗 正在访问教务处...');
    await page.goto('http://zhjw.scu.edu.cn/loginAction.do');

    console.log('📸 正在截图...');
    await page.screenshot({ path: 'login_test.png' });

    console.log('✅ 测试成功！截图已保存为 login_test.png');

    await new Promise(r => setTimeout(r, 5000));
    await browser.close();
})();
