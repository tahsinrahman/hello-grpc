// Code generated by ./hack/browserify.py
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var _ = require('lodash');

var apis = _.merge({},
    require('./apis/hello/v1alpha1/hello.gw.js'),
    require('./apis/status/status.gw.js')
);
module.exports = apis.appscode.hello;
