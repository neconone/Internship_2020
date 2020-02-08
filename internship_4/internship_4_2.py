from flask import Flask, jsonify, request
from urllib.request import urlopen
from bs4 import BeautifulSoup
# Please install library flask, request, beautifulsoup4
app = Flask(__name__)

@app.route('/companies', methods={'GET'})
def index():
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
    data_api = []
    for x in data:
        data_api.append({"logo": "https://theinternship.io/company/"+x[0]})
    return jsonify({"companies": data_api})

if __name__ == '__main__':
    app.run(debug=True)