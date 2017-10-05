'use strict';

console.log('Loading function');

const doc = require('dynamodb-doc');

const dynamo = new doc.DynamoDB();


/**
 * Demonstrates a simple HTTP endpoint using API Gateway. You have full
 * access to the request and response payload, including headers and
 * status code.
 *
 * To scan a DynamoDB table, make a GET request with the TableName as a
 * query string parameter. To put, update, or delete an item, make a POST,
 * PUT, or DELETE request respectively, passing in the payload to the
 * DynamoDB API as a JSON body.
 */
exports.handler = (event, context, callback) => {
    //console.log('Received event:', JSON.stringify(event, null, 2));

    const done = (err, res) => callback(null, {
        statusCode: err ? '400' : '200',
        body: err ? err.message : JSON.stringify(res),
        headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
            //'Access-Control-Allow-Origin': 'therealplato.com'
        },
    });

    switch (event.httpMethod) {
        case 'GET':
            dynamo.scan({ TableName: event.queryStringParameters.TableName }, (err, res) => {
                if(err){
                  return done(err, null)
                }
                var n;
                if (res.Count === 0) {
                  n = 0
                } else {
                  n = res.Items[0].n || 1
                }
                return done(null, {n:n})
            });
            break;
        case 'POST':
            dynamo.scan({ TableName: event.queryStringParameters.TableName }, (err, res) => {
                if(err){
                  return done(err, null)
                }
                var n;
                if (res.Count === 0) {
                  n = 1
                } else {
                  n = (res.Items[0].n || 0) + 1
                }
                dynamo.putItem({TableName:"com-therealplato-counter",Item:{n:n}}, (err, res) =>{
                    if (err) {return done(err, null)}
                    return done(null, {n:n})
                });
            });

            break;


        /*
        case 'PUT':
            dynamo.updateItem(JSON.parse(event.body), done);
            break;
        case 'DELETE':
            dynamo.deleteItem(JSON.parse(event.body), done);
            break;
            */
        default:
            done(new Error(`Unsupported method "${event.httpMethod}"`));
    }
};

