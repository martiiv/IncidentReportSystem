import Table from "../components/Table";
import Persons from "../constants/DummyPersons";
import RowWR from "../constants/RowWR";
import {Component} from "react";
import "./WarningReceiver.css"
import fetchData from "../middleware/FetchData";
import {RECEIVER_URL} from "../constants/WebURL";
import deleteData from "../middleware/deleteData";
import {forEach, map} from "react-bootstrap/ElementChildren";

class WarningReceiver extends Component {
    state = {
        value: [],
        data: []
    };

    /**
     * Function that will get store the id of the checked persons in an array.
     * If unchecked the id will be removed from the array.
     * @param val person to be checked or unchecked.
     */
    onChangeValueHandler = val => {
        if (val.target.checked){
            this.setState(prevState =>( {
                    value: [...prevState.value, val.target.value]
                })
            )
        }else{
            let filteredArray = this.state.value.filter(item => item !== val.target.value)
            this.setState({value: filteredArray});
        }
    };

    async componentDidMount() {
        await fetchData(RECEIVER_URL)
            .then((warningReceiver) => {
                console.log(warningReceiver.data)
                this.setState({data: warningReceiver.data})
            })
    }


    deleteWarningReceivers = async () => {
        const {value} = this.state;
        console.log(value)

        let deleteBody = [
            value.map( val => "{" + val + "},")
        ]

        console.log(deleteBody)



        // eslint-disable-next-line no-restricted-globals
        if (confirm('Are you sure you want to save this thing into the database?')) {
            // Save it!
            console.log(deleteBody)

        } else {
            // Do nothing!
            console.log('Thing was not saved to the database.');
        }


    }



    render() {
         const {value, data} = this.state;
        return (
             <div>
                 <h1>Warning Receiver</h1>
                 <Table
                     type={"WR"}
                     data={data}
                     row={RowWR}
                     value={value}
                     onChangeValue={this.onChangeValueHandler}
                 />
                 <div className={"button-group"} style={{display: "flex", justifyContent: "center"}}>
                     <button className={"btn btnGroup"} onClick={this.deleteWarningReceivers}>Delete</button>
                     <button className={"btn btnGroup"}>Create</button>
                     <button className={"btn btnGroup"}>Group</button>
                 </div>
             </div>
         );
     }
}

export default WarningReceiver;



