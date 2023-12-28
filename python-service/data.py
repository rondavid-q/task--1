from flask import Flask, request, jsonify

app = Flask(__name__)
port = 8001

@app.route('/process-data', methods=['POST'])
def process_data():
    try:
        data = request.get_json()
        sorted_data = sorted(data['numbers'])
        return jsonify(sorted_data)
    except Exception as e:
        return jsonify({'error': str(e)}), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=port)

