from flask import Flask, render_template

app = Flask(__name__)

@app.route('/')
def index():
    services = [
        {
            'title': 'Infant Care',
            'description': 'Specialized care for kids up to 6 weeks, with focus on safety, comfort, and early development.'
        },
        {
            'title': 'Preschool Program',
            'description': 'Engaging educational activities designed through school ages up to 11 years old, including learning, play, and social development.'
        }
    ]
    hours = {
        'time': '5:30am - 6:00pm',
        'days': 'Monday through Friday'
    }
    address = '400 Woodside Dr, Cornell, WI 54732'
    return render_template('index.html', services=services, hours=hours, address=address)

if __name__ == '__main__':
    app.run(debug=True, port=8000)
