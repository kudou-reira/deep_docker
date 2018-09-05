# https://www.peterbe.com/plog/how-to-create-react-app-with-docker
# https://mherman.org/blog/2017/12/07/dockerizing-a-react-app/
# https://stackoverflow.com/questions/47648792/communicating-between-different-docker-services-in-docker-compose

# dockerizing go https://medium.com/@pierangelo1982/dockerize-an-existing-golang-web-server-a0b9cf7746a9
# https://blog.blindside.io/docker-compose-go-c2bb7bed0bab

# communication https://stackoverflow.com/questions/47648792/communicating-between-different-docker-services-in-docker-compose

# https://medium.com/@joatmon08/using-containers-to-learn-nginx-reverse-proxy-6be8ac75a757

# reverse proxying with nginx https://medium.com/@joatmon08/using-containers-to-learn-nginx-reverse-proxy-6be8ac75a757
# nginx api https://dzone.com/articles/configuring-nginx-to-proxy-rest-requests-across-mu
# setting up nginx as a reverse proxy for microservices https://memz.co/reverse-proxy-nginx-docker-microservices/

# change create react act starting port https://tech.amikelive.com/node-830/reactjs-changing-default-port-3000-in-create-react-app/

# have go server manage the database

# recurrent neural network stuff
# A deep learning framework for financial time series using stacked autoencoders and long-short term memory
# http://journals.plos.org/plosone/article?id=10.1371/journal.pone.0180944

# Financial Market Time Series Prediction with RNN
# http://cs229.stanford.edu/proj2012/BernalFokPidaparthi-FinancialMarketTimeSeriesPredictionwithRecurrentNeural.pdf

# Predict Stock Prices using RNN (Lilian Weng)
# https://lilianweng.github.io/lil-log/2017/07/08/predict-stock-prices-using-RNN-part-1.html
# https://lilianweng.github.io/lil-log/2017/07/22/predict-stock-prices-using-RNN-part-2.html


# base image
FROM node:9.6.1

# set working directory
RUN mkdir /usr/src/app
WORKDIR /usr/src/app

# add `/usr/src/app/node_modules/.bin` to $PATH
ENV PATH /usr/src/app/node_modules/.bin:$PATH

# install and cache app dependencies
COPY package.json /usr/src/app/package.json
RUN npm install
RUN npm install react-scripts@1.1.1 -g --silent

# start app
CMD ["npm", "start"]