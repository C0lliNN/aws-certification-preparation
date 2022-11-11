exports.handler = function (event, context) {
  console.info('EVENT\n' + JSON.stringify(event, null, 2));
  return context.logStreamName;
};
