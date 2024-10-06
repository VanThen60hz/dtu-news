import React from "react";
import Tag from "antd/lib/tag";
const statusMap = {
    complete: <Tag color="green">Complete</Tag>,
    inProgress: <Tag color="orange">In Progress</Tag>,
    pending: <Tag color="yellow">Pending</Tag>,
};

export const StatusTag = ({ status }) => statusMap[status];
