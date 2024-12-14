import { createSignal, onMount } from "solid-js";
import { Col, Container, Row } from "solid-bootstrap";
import { SolidApexCharts } from "solid-apexcharts"

import css from './dashboard.module.css'

function Dashboard(props) {
    const options = {
        xaxis: {
            categories: [1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998],
        },
        theme: {
            mode: "dark",
            palette: "palette3",
        },
        chart: {
            background: "#212529",
            toolbar: {
                show: false,
            },
        },
    };
    const [series, setSeries] = createSignal([
        {
            name: 'series-1',
            data: [30, 40, 35, 50, 49, 60, 70, 91],
        },
    ]);
    onMount(() => {
        const max = 90;
        const min = 20;

        setInterval(() => {
            setSeries((prevSeries) => {
                const newData = prevSeries[0].data.map(() => {
                    return Math.floor(Math.random() * (max - min + 1)) + min
                })

                return [{ name: 'series-1', data: newData }]
            });
        }, 1000)
    })
    return (
        <Container>
            <Row>
                <Col sm="auto" md="2" >
                    <Container class={css.box}>
                        heelo<br />
                        testing
                    </Container>
                </Col>
                <Col md="auto">
                    <div class={css.box}>
                        <SolidApexCharts type="bar" options={options} series={series()} />
                    </div>
                </Col>
                <Col sm="auto" md="3">
                    3 of 3
                </Col>
            </Row>
        </Container>

    );
}

export default Dashboard;
