const fastify = require('fastify')({logger: true});

fastify.get('/', (request, reply) => reply.code(400).send({ name: 'Max Mustermann'}));

fastify.get('/products', (request, reply) => {
    const products = ["Milk", "Bread", "Honey", "Potatoes", "Chocolate"];
    reply.send(products);
});

fastify.listen(3001);
