FROM python:3.7

WORKDIR /code

RUN git clone https://github.com/RasaHQ/financial-demo.git financio

WORKDIR financio

RUN pip install --upgrade pip

RUN pip install -r requirements.txt

COPY . .

RUN python -m spacy download en_core_web_md

RUN python -m spacy link en_core_web_md en

RUN rasa train

EXPOSE 5056

CMD [ "/bin/bash" ]