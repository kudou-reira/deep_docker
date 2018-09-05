# Observatory Service

# Import framework
from flask import Flask
from flask_restful import Resource, Api, fields, marshal_with
from flask_cors import CORS

# http://flask-restful.readthedocs.io/en/0.3.5/quickstart.html
# communicate between multiple docker-compose files

# https://stackoverflow.com/questions/26980713/solve-cross-origin-resource-sharing-with-flask

# Instantiate the app
app = Flask(__name__)
CORS(app, origins="http://localhost:3000", allow_headers=[
    "Content-Type", "Authorization", "Access-Control-Allow-Credentials"],
    supports_credentials=True)
api = Api(app)

class Observatory(Resource):
    def get(self):
        return {
            'Galaxies': ['Milkyway', 'Andromeda', 
            'Large Magellanic Cloud (LMC)']
        }

# Create routes
api.add_resource(Observatory, '/')

# Run the application
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3000, debug=True)