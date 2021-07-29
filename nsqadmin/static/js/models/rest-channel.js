var _ = require('underscore');

var AppState = require('../app_state');
var Backbone = require('backbone');

var RestChannel = Backbone.Model.extend({
    idAttribute: 'id',

    constructor: function RestChannel() {
        Backbone.Model.prototype.constructor.apply(this, arguments);
    },

    url: function() {
        return AppState.apiPath('/rest-channels/' + encodeURIComponent(this.get('id')));
    },

    parse: function(response) {
        response['nodes'] = _.map(response['nodes'] || [], function(node) {
            var nodeParts = node['node'].split(':');
            var port = nodeParts.pop();
            var address = nodeParts.join(':');
            var hostname = node['hostname'];
            node['show_broadcast_address'] = hostname.toLowerCase() !== address.toLowerCase();
            node['hostname_port'] = hostname + ':' + port;
            return node;
        });
        return response;
    }
});

module.exports = RestChannel;
