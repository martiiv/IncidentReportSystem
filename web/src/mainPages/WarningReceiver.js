import Table from "../components/Table";
import RowWR from "../constants/RowWR";
import {Component} from "react";
import "./WarningReceiver.css"
import fetchData from "../middleware/FetchData";
import {RECEIVER_URL} from "../constants/WebURL";
import deleteData from "../middleware/deleteData";
import {Link} from "react-router-dom";
import {ReactNotifications, Store} from "react-notifications-component";


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
                    value: [...prevState.value, {id: (val.target.value)}]
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
        // eslint-disable-next-line no-restricted-globals
        if (confirm('Are you sure you want to save this thing into the database?')) {
            await deleteData(RECEIVER_URL, value)
                .then(() => {
                    Store.addNotification({
                        title: "Successfully Deleted!",
                        message: "Warning Receiver is successfully deleted",
                        type: "success",
                        insert: "top",
                        container: "top-right",
                        animationIn: ["animate__animated", "animate__fadeIn"],
                        animationOut: ["animate__animated", "animate__fadeOut"],
                        dismiss: {
                            duration: 2000,
                            onScreen: true
                        }
                    });
                })
                .then(setTimeout(() => {
                    window.location.reload(true)
                }, 2000))
                .catch(() => {
                    alert("something went wrong. please try again")
                })

        }
    }


    render() {
        const {value, data} = this.state;
        return (
            <div>
                <h1>Warning Receiver</h1>
                <ReactNotifications />

                <Table
                    type={"WR"}
                    data={data}
                    row={RowWR}
                    value={value}
                    onChangeValue={this.onChangeValueHandler}
                />
                <div className={"button-group"} style={{display: "flex", justifyContent: "center"}}>
                    <button className={"btn btnGroup"} onClick={this.deleteWarningReceivers}>Delete</button>
                    <Link to={RECEIVER_URL + "/new"}>
                        <button className={"btn btnGroup"}>Create</button>
                    </Link>

                    <Link to={RECEIVER_URL + "/group"}>
                        <button className={"btn btnGroup"}>Group</button>
                    </Link>
                </div>
            </div>
        );
    }
}

export default WarningReceiver;



