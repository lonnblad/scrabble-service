const AWS = require('aws-sdk');

exports.handler = (event) => {
  const stepfunctions = new AWS.StepFunctions();

  const params = {
    stateMachineArn: process.env.STATE_MACHINE_ARN,
    input: JSON.stringify(event)
  };

  stepfunctions.startExecution(params).promise()
    .then(data => { console.log({ data }); })
    .catch(err => { console.log({ err }); });
};