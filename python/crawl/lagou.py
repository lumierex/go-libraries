import threading
from queue import Queue
import requests
import json
import time
import pdfkit


class LaGou_Article_Spider():
    def __init__(self, url):
        self.url = url
        self.headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36',
            # 'Cookie': 'smidV2=202101111726286fdab40363e33b94f72ffce351360830006edcb51a74fa590; sensorsdata2015session={}; EDUJSESSIONID=ABAAAECABCAAACD05B4F2E2ACEB924F8173DE43AEADEF3D; thirdDeviceIdInfo=[{"channel":1,"thirdDeviceId":"WHJMrwNw1k/FmnQSb5wOPU8wqIM5PRmQUGAdmgqrco8fI1UpDCinY3Sm6BO99jiYUv+j/K2n33QW4LTFLRGSXnV21pN9QZSbMdCW1tldyDzmQI99+chXEitFpV+4gsU+z9lCUKKcsmkTaFO8webhNijYmmmXo8LlTkQE5YcNLqNriNYPfoOP/bkHt+iuM33fO73TKvvAESRl6cv0xerXfAe8Vg+5pExYXHm13s1EruLixF+BfZKLGhypuVy3QpiNySBMZJRjFRB4=1487582755342"},{"channel":2,"thirdDeviceId":"140#l5FoVm5EzzPCiQo23z7z4pN8s7aV/dNZWMSvwTfsc5JBao5qgrWT0YnHcpSCay/XIVTLt3hqzznTb+GhzFzzzjVbbYJqlQzx2DD3VthqzFcJLX8+lpYazDrbV25cDF+9ONdOHaU+PkICJJGdMzv60GmIoR+25szLvQPa8HCUMIlAnk9OTH0OQ8P0yaJhdL6LyBg3HILMbX5jvOoe5kCPOCD4ATnhx+NOeWZzHyNM/DK55bmIiZTXJvPejD5U5P+eSDLQSgCRPuB4Ujy0fzRyQmVuR3IYW+FaE1K5XdfAoq5tWGnjZxDD15UpR11Ye5TgmYHWLDZywxdBkhNcpRXr0e0urmIuZR9keAWz2dGrOTrcEHKVyjJZv0O292rrecnYk7cObDKASJtPiMZLzM+Bi23NYejAc5bYLPTTJerVdQqRCXLPFmBX3tHIMK60jUiCOwLSiC1/ghNN0AGP+6PsBcsIWfFN3VaJ9Rj9Vutwz1AzMYOEDuhi3TMSs3cdS2qnVK6QJUM7VfPVcjr0KzXM+dZKeC5eKygTut37wM4tGfKwmk9fE4SK6CTYZfz8t5ZLKAYOkwgnI27Yw1ZUEczM6PBrNjtmxAqOcyK+l1u1aV8BYVJZ8FdVOF0GadjmQMNyNUzPn4BbPJsjWSZlvIMIXkThFC/bRdrmr96EFawiiGD9YDd38TZK3QQLU7Dh4fCQAJeQ/WhTebRuPP/BKm3LftDeC1PqGsFpzbo+51otilWed+syE5rIzNJ1Uw/dfjHC5ogDNBL0f+7elV941VTumIPMKmx4IifAxFXg8T2PGKzyYdDD+jRGEyQyajpAWblevfBPWHO9mAIc1RY8RJbsr/3801t+1P05sGbLxzp9+IhBhTXiH99ww7HKI8cS7XCyRr8FZ3lmUbaoDyATy5PR3y9T4xTcuQqV41SuKoMvas3VUxjyKPnmzPCF2YpiSXGre/WnEb==,undefined"}]; user-finger=761bab24aa2b86b0a8a0cd0c04f0c358; LG_LOGIN_USER_ID=a67c5a847b328ee4247ef639a5e908d132df236e11cec426; LG_HAS_LOGIN=1; _putrc=E30E2B97F4BFC515; login=true; unick=陈; kw_login_authToken="U8LowYzWb4XWQFc/QaIL/1CzmxcLbsY4IwD7MFzlDznNXNVsShq9UrQNrjQXZTBDry70aM6QlkNhRWkF0tAJDkLWt4/5nQGOaxmnAxnVp49QHEvJRVvzLg3KRAZhEy4VbsFVipOLXpZAe5sCxMmqnxCkBDY5s+UnJvNzMKaaKaF4rucJXOpldXhUiavxhcCELWDotJ+bmNVwmAvQCptcy5e7czUcjiQC32Lco44BMYXrQ+AIOfEccJKHpj0vJ+ngq/27aqj1hWq8tEPFFjdnxMSfKgAnjbIEAX3F9CIW8BSiMHYmPBt7FDDY0CCVFICHr2dp5gQVGvhfbqg7VzvNsw=="; gate_login_token=22fcc8422e303d9e15981f0b533ae574442f709025d21835; X_HTTP_TOKEN=42daf4b72327b2816408237161bf5e71415983ed09; sensorsdata2015jssdkcross={"distinct_id":"7588076","first_id":"1789042e7609ac-0d4cf7faeb21e6-5771e33-2073600-1789042e761800","props":{"$latest_traffic_source_type":"自然搜索流量","$latest_search_keyword":"未取到值","$latest_referrer":"https://www.baidu.com/link","$latest_utm_source":"baidujava","$latest_utm_medium":"sspc","$latest_utm_term":"java20913","$os":"Windows","$browser":"Chrome","$browser_version":"89.0.4389.72"},"$device_id":"176f0c4a7fcdd-0b3866cef6271e-f7d123e-2073600-176f0c4a7fd992"}; JSESSIONID=F8958EB110A0359985055831B04B38BE',
            'Cookie': 'smidV2=202101111726286fdab40363e33b94f72ffce351360830006edcb51a74fa590; _bl_uid=1nkdgjnjs2hd6v4ppys93y0pLtF9; sensorsdata2015session=%7B%7D; WEBTJ-ID=2021042%E4%B8%8A%E5%8D%889:46:46094646-1789042e62f271-0c1e4a01478238-5771e33-2073600-1789042e63071f; thirdDeviceIdInfo=%5B%7B%22channel%22%3A1%2C%22thirdDeviceId%22%3A%22WHJMrwNw1k/FmnQSb5wOPU8wqIM5PRmQUGAdmgqrco8fI1UpDCinY3Sm6BO99jiYUv+j/K2n33QW4LTFLRGSXnV21pN9QZSbMdCW1tldyDzmQI99+chXEitFpV+4gsU+z9lCUKKcsmkTaFO8webhNijYmmmXo8LlTkQE5YcNLqNriNYPfoOP/bkHt+iuM33fO73TKvvAESRl6cv0xerXfAe8Vg+5pExYXHm13s1EruLixF+BfZKLGhypuVy3QpiNySBMZJRjFRB4%3D1487582755342%22%7D%2C%7B%22channel%22%3A2%2C%22thirdDeviceId%22%3A%22140%23l5FoVm5EzzPCiQo23z7z4pN8s7aV/dNZWMSvwTfsc5JBao5qgrWT0YnHcpSCay/XIVTLt3hqzznTb+GhzFzzzjVbbYJqlQzx2DD3VthqzFcJLX8+lpYazDrbV25cDF+9ONdOHaU+PkICJJGdMzv60GmIoR+25szLvQPa8HCUMIlAnk9OTH0OQ8P0yaJhdL6LyBg3HILMbX5jvOoe5kCPOCD4ATnhx+NOeWZzHyNM/DK55bmIiZTXJvPejD5U5P+eSDLQSgCRPuB4Ujy0fzRyQmVuR3IYW+FaE1K5XdfAoq5tWGnjZxDD15UpR11Ye5TgmYHWLDZywxdBkhNcpRXr0e0urmIuZR9keAWz2dGrOTrcEHKVyjJZv0O292rrecnYk7cObDKASJtPiMZLzM+Bi23NYejAc5bYLPTTJerVdQqRCXLPFmBX3tHIMK60jUiCOwLSiC1/ghNN0AGP+6PsBcsIWfFN3VaJ9Rj9Vutwz1AzMYOEDuhi3TMSs3cdS2qnVK6QJUM7VfPVcjr0KzXM+dZKeC5eKygTut37wM4tGfKwmk9fE4SK6CTYZfz8t5ZLKAYOkwgnI27Yw1ZUEczM6PBrNjtmxAqOcyK+l1u1aV8BYVJZ8FdVOF0GadjmQMNyNUzPn4BbPJsjWSZlvIMIXkThFC/bRdrmr96EFawiiGD9YDd38TZK3QQLU7Dh4fCQAJeQ/WhTebRuPP/BKm3LftDeC1PqGsFpzbo+51otilWed+syE5rIzNJ1Uw/dfjHC5ogDNBL0f+7elV941VTumIPMKmx4IifAxFXg8T2PGKzyYdDD+jRGEyQyajpAWblevfBPWHO9mAIc1RY8RJbsr/3801t+1P05sGbLxzp9+IhBhTXiH99ww7HKI8cS7XCyRr8FZ3lmUbaoDyATy5PR3y9T4xTcuQqV41SuKoMvas3VUxjyKPnmzPCF2YpiSXGre/WnEb%3D%3D%2Cundefined%22%7D%5D; user-finger=761bab24aa2b86b0a8a0cd0c04f0c358; LG_LOGIN_USER_ID=a67c5a847b328ee4247ef639a5e908d132df236e11cec426; LG_HAS_LOGIN=1; _putrc=E30E2B97F4BFC515; login=true; unick=%E9%99%88; privacyPolicyPopup=false; gate_login_token=22fcc8422e303d9e15981f0b533ae574442f709025d21835; X_HTTP_TOKEN=42daf4b72327b2814371537161bf5e71415983ed09; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%227588076%22%2C%22first_id%22%3A%221789042e7609ac-0d4cf7faeb21e6-5771e33-2073600-1789042e761800%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%2C%22%24latest_utm_source%22%3A%22baidujava%22%2C%22%24latest_utm_medium%22%3A%22sspc%22%2C%22%24latest_utm_term%22%3A%22java20913%22%2C%22%24os%22%3A%22Windows%22%2C%22%24browser%22%3A%22Chrome%22%2C%22%24browser_version%22%3A%2289.0.4389.72%22%7D%2C%22%24device_id%22%3A%22176f0c4a7fcdd-0b3866cef6271e-f7d123e-2073600-176f0c4a7fd992%22%7D',
            'Referer': 'https://kaiwu.lagou.com/course/courseInfo.htm?courseId=17',
            'Origin': 'https://kaiwu.lagou.com',
            'accept': '*/*',
            'Authorization': '22fcc8422e303d9e15981f0b533ae574442f709025d21835',
            'Sec-fetch-dest': 'empty',
            'Sec-fetch-mode': 'cors',
            'Sec-fetch-site': 'same-site',
            'x-l-req-header': '{deviceType:1}'}
        # 发现课程文章html的请求url前面都是一样的最后的id不同而已
        self.textUrl = 'https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessonDetail?lessonId='
        self.queue = Queue()  # 初始化一个队列
        self.error_queue = Queue()

    def parse_one(self):
        """
        :return:获取文章html的url
        """
        # id_list=[]
        html = requests.get(url=self.url, headers=self.headers).text

        dit_message = json.loads(html)
        message_list = dit_message['content']['courseSectionList']
        # print(message_list)
        for message in message_list:
            for i in message['courseLessons']:
                true_url = self.textUrl+str(i['id'])
                self.queue.put(true_url)  # 文章的请求url
        return self.queue

    def get_html(self, true_url):
        global article_name
        html = requests.get(url=true_url, timeout=10,
                            headers=self.headers).text
        dit_message = json.loads(html)
        str_html = str(dit_message['content']['textContent'])
        article_name1 = dit_message['content']['theme']
        if "|" or '?' or '/' in article_name1:
            article_name = article_name1.replace("|" and '?' and '/', "-")
        else:
            article_name = article_name1
        self.htmltopdf(str_html, article_name)

    def htmltopdf(self, str_html, article_name):
        # C:\Program Files\wkhtmltopdf\bin
        path_wk = r'C:\Program Files\wkhtmltopdf\bin\wkhtmltopdf.exe'
        config = pdfkit.configuration(wkhtmltopdf=path_wk)
        options = {
            'page-size': 'Letter',
            'encoding': 'UTF-8',
            'custom-header': [('Accept-Encoding', 'gzip')]
        }
        pdfkit.from_string(str_html, "D:\\data\\{}.pdf".format(
            article_name), configuration=config, options=options)

    def thread_method(self, method, value):  # 创建线程方法
        thread = threading.Thread(target=method, args=value)
        return thread

    def main(self):

        thread_list = []
        true_url = self.parse_one()
        print("true_url: ", true_url)
        while true_url is not None and not true_url.empty():
            for i in range(10):  # 创建线程并启动
                if not true_url.empty():
                    m3u8 = true_url.get()
                    print(m3u8)
                    thread = self.thread_method(self.get_html, (m3u8,))
                    thread.start()
                    print(thread.getName() + '启动成功,{}'.format(m3u8))
                    thread_list.append(thread)
                else:
                    break
            while len(thread_list) != 0:
                for k in thread_list:
                    k.join()  # 回收线程
                    print('{}线程回收完毕'.format(k))
                    thread_list.remove(k)


if __name__ == "__main__":
    #  https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessons?courseId=46

    run = LaGou_Article_Spider(
        " https://gate.lagou.com/v1/neirong/kaiwu/getCourseLessons?courseId=46")
    run.main()