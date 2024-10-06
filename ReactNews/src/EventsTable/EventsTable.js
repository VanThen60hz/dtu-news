import React from "react";

import { actionService } from "./eventService";
import { StatusTag } from "../StatusTag";
import { Flex, Table, Button } from "antd";

const EventsTable = ({ eventsData }) => {
    const tableColumns = [
        {
            title: "ID",
            dataIndex: "id",
            key: "id",
        },
        {
            title: "Title",
            dataIndex: "title",
            key: "title",
        },
        {
            title: "Summary",
            dataIndex: "summary",
            key: "summary",
        },
        {
            title: "Content",
            dataIndex: "content",
            key: "content",
        },
        {
            title: "Action",
            key: "action",
            width: 210,
            render: (text, record) => (
                <>
                    <Flex justify={"space-around"} align={"center"}>
                        <Button
                            type="default"
                            style={{ borderColor: "#1890ff", color: "#1890ff" }}
                            onClick={() => handleAction(record)}
                        >
                            Update
                        </Button>
                        <Button
                            type="default"
                            onClick={() => handleAnotherAction(record)}
                            style={{ marginLeft: "auto" }}
                            danger
                        >
                            Delete
                        </Button>
                    </Flex>
                </>
            ),
        },
    ];

    const handleAction = (currentEvent) => {
        actionService(currentEvent);
    };

    const handleAnotherAction = (currentEvent) => {
        actionService(currentEvent);
    };

    return <Table dataSource={eventsData} columns={tableColumns} bordered rowKey="id" />;
};

export { EventsTable };
