const fastify = require('fastify')({logger: true});

fastify.get('/', (request, reply) => reply.code(400).send({ name: 'Max Mustermann'}))

fastify.listen(3001);
