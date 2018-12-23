function drawLine(el, dataTable) {
    $("#" + el).html('');
    cols = [];
    labels = [];
    visibleCols = dataTable.bootstrapTable('getVisibleColumns');
    for (i in visibleCols) {
        if (visibleCols[i].field !== "note_time" && visibleCols[i].field !== undefined) {
            cols.push(visibleCols[i].field);
            labels.push(visibleCols[i].title);
        }
    }
    colors = ['#afd8f8', '#edc240', '#cb4b4b', '#9440ed', '#32C2CD', '#gray', '#6BF9D1', '#D2DA38', '#00FF77', '#FF00AB'];
    Morris.Line({
        element: el,
        data: dataTable.bootstrapTable('getData'),
        xkey: 'note_time',
        ykeys: cols,
        labels: labels,
        lineWidth: 2,
        pointSize: 2,
        fillOpacity: 0,
        hideHover: 'auto',
        behaveLikeLine: true,
        resize: true,
        pointStrokeColors: colors,
        lineColors: colors,
    });
}

function darwArea(el, dataTable, type) {
    $("#" + el).html('');
    tableData = dataTable.bootstrapTable('getData');
    showData = [];
    labels = [];
    time = null;
    for (i in tableData) {
        name = tableData[i].name;
        if (type === 'times') {
            value = tableData[i].times;
        }
        else {
            value = tableData[i].poundage;
        }
        if (time !== tableData[i].note_time) {
            time = tableData[i].note_time;
            data = [];
            data['note_time'] = time;
            data[name] = value;
            if (time !== undefined) {
                showData.push(data)
            }
        } else {
            data[name] = value;
        }
        if (!labels.exists(name) && name.length > 0) {
            labels.push(name)
        }
    }
    Morris.Area({
        element: el,
        data: showData,
        xkey: 'note_time',
        ykeys: labels,
        labels: labels,
        pointSize: 1,
        hideHover: 'auto',
        resize: true
    });
}

function darwDonut(el, dataTable, type) {
    $("#" + el).html('');
    tableData = dataTable.bootstrapTable('getData');
    showData = [];
    tmpdata = {};
    for (i in tableData) {
        if (type === 'times') {
            value = tableData[i].times;
        }
        else {
            value = tableData[i].poundage;
        }
        name = tableData[i].name;
        if (tmpdata[name] === undefined) {
            tmpdata[name] = 0
        }
        tmpdata[name] += value;
    }
    for (key in tmpdata) {
        data = [];
        data['label'] = key;
        data['value'] = tmpdata[key];
        if (tmpdata[key] > 0) {
            showData.push(data);
        }
    }
    Morris.Donut({
        element: el,
        data: showData,
        colors: ['#afd8f8', '#edc240', '#cb4b4b', '#9440ed', '#32C2CD', '#6BF9D1', '#D2DA38', '#00FF77', '#FF00AB'],
        resize: true
    });
}