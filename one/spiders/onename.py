# -*- coding: utf-8 -*-
import scrapy

from one.items import OneItem


class OnenameSpider(scrapy.Spider):
    name = 'onename'
    allowed_domains = ['m.wufazhuce.com']
    start_urls = ['http://m.wufazhuce.com/one/2174']

    def parse(self, response):
        item = OneItem()

        item["imgUrl"] = response.xpath("//*[@id='myPage_angOne_view']/div[4]/div[1]/img/@src").extract()
        item["text"] = response.xpath("//*[@id='myPage_angOne_view']/div[4]/div[1]/p[4]/text()").extract()
        item["day"] = response.xpath("//*[@id='myPage_angOne_view']/div[4]/div[1]/p[2]/text()").extract()
        item["month"] = response.xpath("//*[@id='myPage_angOne_view']/div[4]/div[1]/p[3]/text()").extract()
        str = response.url
        name = str.split("/")
        item["find"] = name[len(name)-1]
        print(item["find"])
        yield item

        for i in range(2170, 3000):
            url = "http://m.wufazhuce.com/one/%d" % (i)
            print(url + "\n")

            yield scrapy.Request(url, callback=self.parse)

    # def get_http_status(self, response):
    #     if response.status != 404:
    #         yield scrapy.Request(response.url, callback=self.parse)
    #     else:
    #         print("404 界面 " + response.url)
