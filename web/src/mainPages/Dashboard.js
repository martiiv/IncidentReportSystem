import Table from "../components/Table"
import RowName from "../constants/RowName";
import {useEffect, useState} from "react";
import fetchData from "../middleware/FetchData";
import {INCIDENT_URL} from "../constants/WebURL";


/**
 * Dashboard page, that will display the incident log.
 *
 * @returns {JSX.Element}
 * @constructor
 */
function Dashboard() {
    const [incident, setIncident] = useState([])

    /**
     * Function that will be called when the page is loaded.
     * Function will fetch incident data.
     */
    useEffect( () => {
        const fetch = async () => {
            const data = await fetchData(INCIDENT_URL);
            setIncident(data)
        }
        fetch()
            // make sure to catch any error
            .catch(() => alert("No connection, try again later"));
    }, [])


    /**
     * If data is not fetched, display "Loading"
     */
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
