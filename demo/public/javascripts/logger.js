/**
 * Created by liqiao on 8/14/15.
 */

var log = document.createElement('div');
log.setAttribute('id', 'log');
document.body.appendChild(log);

var logger = {
    i: function(info) {
        add(info, 'i');
    },
    d: function(debug) {
        add(debug, 'd');
    },
    e: function(err) {
        add(err, 'e');
    }
};

function add(msg, level) {
    var row = document.createElement('div');
    row.setAttribute('class', 'log-row log-' + level);
    row.innerHTML = "\n"+msg;
    document.querySelector('#log').appendChild(row);
}