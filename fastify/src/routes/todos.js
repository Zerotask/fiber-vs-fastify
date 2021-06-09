const basePath = "/todos";

const AddTodos = (fastify) => {
    fastify.get(basePath + "/", (request, reply) => reply.code(400).send("Todos get test"))
}

module.exports = { AddTodos };