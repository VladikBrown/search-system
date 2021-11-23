import './Results.css';

export default function Results(props) {
    const {items} = props;

    return (<div className="ItemList">
            {items.map((item, i) => <div key={i} className="Item">{item.name}</div>)}
    </div>)
}