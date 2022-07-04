from selenium import webdriver
import time


def selenium_get_cookies(url='https://www.douyin.com'):
    """无头模式提取目标链接对应的cookie，代码作者：小小明-代码实体"""
    start_time = time.time()
    option = webdriver.ChromeOptions()
    option.add_argument("--headless")
    option.add_experimental_option('excludeSwitches', ['enable-automation'])
    option.add_experimental_option('useAutomationExtension', False)
    option.add_argument(
        'user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36')
    option.add_argument("--disable-blink-features=AutomationControlled")
    print("打开无头游览器...")
    browser = webdriver.Chrome(options=option)
    print(f"访问{url} ...")
    browser.get(url)
    cookie_list = browser.get_cookies()
    # 关闭浏览器
    browser.close()
    cost_time = time.time() - start_time
    print(f"无头游览器获取cookie耗时：{cost_time:0.2f} 秒")
    return {row["name"]: row["value"] for row in cookie_list}


print(selenium_get_cookies("https://www.douyin.com"))