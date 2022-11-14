import Table from "../components/Table"
import RowName from "../constants/RowName";
import {useEffect, useState} from "react";
import fetchData from "../middleware/FetchData";
import {INCIDENT_URL} from "../constants/WebURL";

function Dashboard() {
    const [incident, setIncident] = useState([])

    useEffect( () => {
        const fetch = async () => {
            const data = await fetchData(INCIDENT_URL);
            setIncident(data)
        }

        fetch()
            // make sure to catch any error
            .catch(console.error);
    }, [])


    console.log(incident.data)

    if (incident.data === undefined){
        return <h1>Loading</h1>
    }else {
        return (

        <div className="App" >
            <h1 style={{display: "flex" , justifyContent: "center" }}>Incident Log</h1>

            <Table
                type = {"IL"}
                data={incident.data}
                row={RowName}
            />
        </div>
    );}

}

export default Dashboard;
