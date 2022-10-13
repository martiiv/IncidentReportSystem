import Table from "../components/Table"
import Incident from "../constants/DummyData"
import RowName from "../constants/RowName";


function Dashboard() {
    return (
        <div className="App" >
           <h1 style={{display: "flex" , justifyContent: "center" }}>Incident Log</h1>
            <Table
                type = {"IL"}
                data={Incident}
                    row={RowName}
            />
        </div>
    );
}

export default Dashboard;
