resource "aws_ecs_cluster" "main" {
  name = "blockchain-client-cluster"
}

resource "aws_ecs_task_definition" "app" {
  family                   = "blockchain-client-task"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "256"
  memory                   = "512"

  container_definitions = jsonencode([{
    name      = "blockchain-client"
    image     = var.container_image
    essential = true
    portMappings = [{
      containerPort = 8080
      hostPort      = 8080
    }]
  }])
}

resource "aws_ecs_service" "app" {
  name            = "blockchain-client-service"
  cluster         = aws_ecs_cluster.main.id
  task_definition = aws_ecs_task_definition.app.arn
  launch_type     = "FARGATE"
  desired_count   = 1

  network_configuration {
    subnets         = [aws_subnet.public.id]
    security_groups = [aws_security_group.ecs_sg.id]
    assign_public_ip = true
  }

  depends_on = [aws_internet_gateway.igw]
}