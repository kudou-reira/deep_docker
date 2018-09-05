# https://hub.docker.com/r/library/python/tags/
FROM python:3.6

# Coping source in current directory into the image
# python:3-onbuild expects the source in /usr/src/app
COPY requirements.txt /usr/src/app/requirements.txt
WORKDIR /usr/src/app
RUN python -m pip install --upgrade pip
RUN pip install -r requirements.txt
COPY . /usr/src/app/

# Commands in a list
CMD ["python", "api.py"]