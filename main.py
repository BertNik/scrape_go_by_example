from bs4 import BeautifulSoup
import requests

class Main:
    links = False
    def __init__(self, url):
        if url == '':
            pass
        else:
            print(url)
            self.base_url = url
            self.response = requests.get(self.base_url)
            self.body = self.response.content
            self.bs_obj()
            self.getLi()
            self.get_a()

    def bs_obj(self):
        self.soup = BeautifulSoup(self.body, 'html.parser')

    def printLinks(self):
        for list_items in self.soup.find_all('li'):
            for a in list_items.find_all('a'):
                print(a['href']) 

    def getLi(self):
        self.li = self.soup.find_all('li')
        
    def get_a(self):
        self.a = self.soup.select('li > a')
        return self.a

    def getLinks(self):
        links = []
        for i in self.soup.select('li > a'):
            print(i['href'])
            links.append(i['href'])
        self.links = links
        return self.links

    def printLinksWithBaseUrl(self):
        if not self.links:
            self.getLinks()
        for link in self.links:
            print("{0}{1}".format(self.base_url, link))

    def getLinksWithBaseUrl(self):
        links_with_base_url = []
        if not self.links:
            self.getLinks()
        for link in self.links:
            links_with_base_url.append("{0}{1}".format(self.base_url, link))
        self.links_with_base_url = links_with_base_url
        return self.links_with_base_url

    def getContentByPath(self,url):
        response = requests.get(url)
        return BeautifulSoup(response.content,'html.parser')

    def getCodeContent(self, url):
        bs = self.getContentByPath(url)
        text_array = []
        for i in bs.select('.leading, .code'):
            text_array.append(i.text)
        return "".join(text_array)
    
    def run(self, main):
        import importlib
        importlib.reload(main)
        m = main.Main('https://gobyexample.com/')
        return m

    def writeFile(self):
        lwiurl = self.getLinksWithBaseUrl()
        for i in lwiurl:
            contents = self.getCodeContent(i)
            if contents is not None:
                f = open("./output/{0}{1}".format(i.split("/")[-1], ".go"), 'w+')
                f.write(contents)
                f.close()
                print( "{1}\n{0}\n{1}".format(contents[0:100], "-"*100))

'''
$ python
>>> import main
>>> m = main.Main('')
>>> _m = m.run(main)
https://gobyexample.com/
>>> _m.
_m.a                       _m.getCodeContent(         _m.getLinksWithBaseUrl(    _m.printLinks(             _m.soup(
_m.base_url                _m.getContentByPath(       _m.get_a(                  _m.printLinksWithBaseUrl(  _m.writeFile(
_m.body                    _m.getLi(                  _m.li                      _m.response                
_m.bs_obj(                 _m.getLinks(               _m.links                   _m.run(                    
>>> _m.printLinksWithBaseUrl()
https://gobyexample.com/hello-world
https://gobyexample.com/values
https://gobyexample.com/variables
https://gobyexample.com/constants
https://gobyexample.com/for
https://gobyexample.com/if-else
https://gobyexample.com/switch
https://gobyexample.com/arrays
https://gobyexample.com/slices
https://gobyexample.com/maps
https://gobyexample.com/range
https://gobyexample.com/functions
https://gobyexample.com/multiple-return-values
https://gobyexample.com/variadic-functions
https://gobyexample.com/closures
...
>>> _m.writeFile()
----------------------------------------------------------------------------------------------------




package main


import "fmt"


$ go run hello-world.go
hello world


$ go build hello-world.go

#hot reload module if changes are made to main.py
_m = m.run(main)
'''