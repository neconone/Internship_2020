from urllib.request import urlopen
from bs4 import BeautifulSoup

url = "https://theinternship.io/"
html = urlopen(url)
soup = BeautifulSoup(html, 'lxml')
companies = soup.find_all("div", {"class": "partner"})

data = []
for div_tag in companies:
    img_tag = div_tag.find('img', {"class": "center-logos"})['src']
    span_tag = div_tag.find('span', {"class": "list-company"}).text
    data.append((img_tag, span_tag))

def sortSecond(val):
    return len(val[1])
data.sort(key=sortSecond, reverse=False)
for x in data:
    print(x[0])

