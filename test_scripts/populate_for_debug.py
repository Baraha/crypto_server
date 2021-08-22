import pip
pip.main(['install', '-U', 'mongoengine'])

from mongoengine import *


class Cryptocurrency(Document):
    coin_id = StringField()
    rank = StringField(max_length=50)
    interval = IntField()
    priceUsd = StringField()
    isHandle = BooleanField(default=False)

if __name__ == "__main__":
    connect(host="mongodb://127.0.0.1:27017/test")
    data = Cryptocurrency(coin_id='bitcoin', interval=30).save()
    print(data.pk)


