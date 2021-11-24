import './Metrics.css';
import Table from "rc-table";
import {CartesianGrid, Line, LineChart, XAxis, YAxis} from 'recharts';

function Metrics(props) {
    const {metrics} = props;
    const {docSetMetrics, docSeqMetrics, accuracyGraph} = metrics;
    return (
        <div>
            <DocSetMetrics docSetMetrics={docSetMetrics}/>
            <DocSeqMetrics docSeqMetrics={docSeqMetrics}/>
            <AccuracyGraph accuracyGraph={accuracyGraph}/>
        </div>
    )
}

function DocSetMetrics(props) {
    const columns = [
        {
            title: 'Accuracy',
            dataIndex: 'accuracy',
            key: 'accuracy',
            width: 200,
        },
        {
            title: 'Error',
            dataIndex: 'error',
            key: 'error',
            width: 200,
        },
        {
            title: 'fMeasure',
            dataIndex: 'fMeasure',
            key: 'fMeasure',
            width: 200,
        },
        {
            title: 'Precision',
            dataIndex: 'precision',
            key: 'precision',
            width: 200,
        },
        {
            title: 'Recall',
            dataIndex: 'recall',
            key: 'recall',
            width: 200,
        },
    ];
    const {docSetMetrics} = props;
    docSetMetrics.key = 1;
    return (
        <div>
            <h1>Doc Set Metrics</h1>
            <Table
                style={{width: 1000}}
                className="Table"
                rowClassName="TableRow"
                columns={columns}
                data={[docSetMetrics]}
                useFixedHeader={true}
                tableLayout="auto"
            />
        </div>
    )
}

function DocSeqMetrics(props) {
    const columns = [
        {
            title: 'AveragePrecision',
            dataIndex: 'averagePrecision',
            key: 'averagePrecision',
            width: 200,
        },
        {
            title: 'Precision',
            dataIndex: 'precision',
            key: 'precision',
            width: 200,
        },
        {
            title: 'rPrecision',
            dataIndex: 'rPrecision',
            key: 'rPrecision',
            width: 200,
        },
    ];
    const {docSeqMetrics} = props;
    docSeqMetrics.key = 1;
    return (
        <div>
            <h1>Doc Seq Metrics</h1>
            <Table
                style={{width: 600}}
                className="Table"
                rowClassName="TableRow"
                columns={columns}
                data={[docSeqMetrics]}
                useFixedHeader={true}
                tableLayout="auto"
            />
        </div>
    )
}

function AccuracyGraph(props) {
    const {accuracyGraph} = props;
    console.log(accuracyGraph)
    return (
        <div className="Chart">
            <h1>Accuracy graph</h1>
            <LineChart width={600} height={300} data={accuracyGraph.points}>
                <Line type="monotone" dataKey="precision" stroke="#8884d8"/>
                <CartesianGrid stroke="#ccc"/>
                <XAxis dataKey="recall"/>
                <YAxis dataKey="precision"/>
            </LineChart>
        </div>
    )
}

export default Metrics;