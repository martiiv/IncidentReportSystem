import { useState } from 'react';
import "./IncidentReporting.css"

function IncidentReportingForm() {
    const [state, setState] = useState({
        name: "",
        context: "",
        information: "",
        receiver: "",
        countermeasure: ""
    })

    const [isPending, setIsPending] = useState(false);
    const handleSubmit = async (event) => {
        event.preventDefault();
        setIsPending(true)
         fetch('http://localhost:8080/incident/sendMail/', ({
            method: 'POST',
            headers: {
                    'Content-type': 'application/json; charset=UTF-8',
            },
            body: JSON.stringify(state)
        }))
            .then((response) => response.json())
            .then((json) => {
                console.log(json);
                setIsPending(false)
            })
             .catch((err) => {
             console.log(err.message);
         });
    }

    function handleChange(evt) {
        const value = evt.target.value;
        setState({
            ...state, [evt.target.name]: value
        })
    }

    return (
        <div className={"incidentCreate"}>
            <h2>New Incident</h2>
            <form onSubmit={handleSubmit}>
                <label>Incident Title:
                    <input
                        type={"text"}
                        name={"name"}
                        required
                        value={state.name}
                        onChange={handleChange}
                    />
                </label>

                <label>Tags:
                    <input
                        type={"text"}
                        name={"context"}
                        value={state.context}
                        onChange={handleChange}
                    />
                </label>

                <label>Receiver:
                    <input
                        type={"text"}
                        name={"receiver"}
                        required
                        value={state.receiver}
                        onChange={handleChange}
                    />
                </label>

                <label>Incident description:
                    <textarea
                        name={"information"}
                        required
                        value={state.information}
                        onChange={handleChange}
                    />
                </label>
                <label>Countermeasures</label>
                <input
                    type={"text"}
                    name={"countermeasure"}
                    value={state.countermeasure}
                    onChange={handleChange}
                />

                { !isPending && <button>SEND INCIDENT</button> }
                { isPending && <button disabled>Sending mail...</button> }
            </form>
        </div>
    )
}

export default IncidentReportingForm
