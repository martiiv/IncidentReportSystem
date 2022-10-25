import {Link} from "react-router-dom";
import {INCIDENT_URL} from "../constants/WebURL";
import {Component} from "react";


class Table extends Component {

    render() {
        const { onChangeValue, type, data, row } = this.props;
        return (
            <div className={"table-div"} style={{display: "flex", justifyContent: "center"}}>
                <table id={"table"}>
                    <tr>
                        {row.map(item =>
                            <th key={item}>
                                {item}
                            </th>
                        )}
                    </tr>

                    {data.map(item => <tr>

                        {type === "WR" &&
                            <td>
                                <label>
                                    <input type="checkbox" value={item.id} onChange={onChangeValue}/>
                                </label>
                            </td>
                        }

                        <td>{type !== "WR" ? item.date : item.name}</td>
                        <td>{item.receiverGroup}</td>
                        <td>{type !== "WR"  ? item.name : item.company}</td>
                        <td>{type === "WR" ? item.email : "Tags"}</td>
                        <td>{type !== "WR" ? item.countermeasure : item.phoneNumber}</td>
                        {type !== "WR" &&
                            <td>
                                <Link to={INCIDENT_URL + "/" + item.id}>
                                    <button value={item.name}>
                                        {item.name}
                                    </button>
                                </Link>
                            </td>
                        }
                    </tr>)
                    }
                </table>
            </div>
        )

    }
}

export default Table
