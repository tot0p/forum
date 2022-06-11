
google.charts.load('current', {'packages':['corechart']});
google.charts.setOnLoadCallback(%s);

function %s() {
    var data = new google.visualization.DataTable();
    data.addColumn('string', 'Topping');
    data.addColumn('number', 'Slices');
    data.addRows(%s);
    var options = {'title':'%s','width':500,'height':500};
    var chart = new google.visualization.PieChart(document.getElementById("%s"));
    chart.draw(data, options);
}