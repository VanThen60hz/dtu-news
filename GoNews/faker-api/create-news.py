import requests
import json
from faker import Faker

fake = Faker()

url = "http://localhost:8000/news"

num_requests = int(input("Number of request: "))

for _ in range(num_requests):
    title = fake.sentence(nb_words=5)
    summary = fake.sentence(nb_words=10)
    content = fake.paragraph(nb_sentences=5)

    payload = json.dumps({
        "title": title,
        "summary": summary,
        "content": content
    })

    headers = {
        'Content-Type': 'application/json'
    }

    response = requests.post(url, headers=headers, data=payload)

    print(response.text)
