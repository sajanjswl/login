data "aws_route53_zone" "zone" {
  name = "${var.dns_zone_name}."
}

resource "aws_route53_record" "app" {
  zone_id = data.aws_route53_zone.zone.zone_id
  name    = "${lookup(var.subdomain, terraform.workspace)}.${data.aws_route53_zone.zone.name}"
  # Canonical name link it to another existing DNS name
  type = "CNAME"
  # time to live,
  ttl = "300"
  #  records can be added to a DNS zone in order to add subdomains to that zone
  records = [aws_lb.api.dns_name]
}

resource "aws_acm_certificate" "cert" {
  # fqdn= fully qualified domain name
  domain_name       = aws_route53_record.app.fqdn
  validation_method = "DNS"

  tags = local.common_tags
  #  This is just to keep our terraform running smooth
  lifecycle {
    create_before_destroy = true
  }
}


# The way Domain validation works is that you assign a DNS record within that domain
# If you assign the random string that a certificate manager gives you to a record in your 
resource "aws_route53_record" "cert_validation" {

  # for_each = {
  #   for dvo in aws_acm_certificate.cert.domain_validation_options : dvo.domain_name => {
  #     name   = dvo.resource_record_name
  #     record = dvo.resource_record_value
  #     type   = dvo.resource_record_type
  #   }
  # }
  # allow_overwrite = true
  # name            = each.value.name
  # records         = [each.value.record]
  # ttl             = 60
  # type            = each.value.type
  # zone_id         = data.aws_route53_zone.zone.zone_id




  # name    = aws_acm_certificate.cert.domain_validation_options.0.resource_record_name
  name = aws_acm_certificate.cert.domain_validation_options.*.resource_record_name[0]
  # type    = aws_acm_certificate.cert.domain_validation_options.0.resource_record_type
  type    = aws_acm_certificate.cert.domain_validation_options.*.resource_record_type[0]
  zone_id = data.aws_route53_zone.zone.zone_id
  records = [
    # aws_acm_certificate.cert.domain_validation_options.0.resource_record_value
    aws_acm_certificate.cert.domain_validation_options.*.resource_record_value[0]
  ]
  ttl = "60"





}

resource "aws_acm_certificate_validation" "cert" {
  certificate_arn         = aws_acm_certificate.cert.arn
  validation_record_fqdns = [aws_route53_record.cert_validation.fqdn]
  # validation_record_fqdns = [for record in aws_route53_record.cert_validation : record.fqdn]
}

# https://github.com/hashicorp/terraform-provider-aws/issues/14447