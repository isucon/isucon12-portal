require 'isuxportal/resources/contestant_instance_pb'

class ContestantInstance < ApplicationRecord
  belongs_to :team

  has_many :benchmark_jobs, inverse_of: :target

  validates :cloud_id, presence: true
  validates :number, presence: true, uniqueness: { scope: :team_id }
  validates :private_ipv4_address, presence: true

  validate :validate_number_in_range

  enum(status: Isuxportal::Proto::Resources::ContestantInstance::Status.descriptor.to_enum.sort_by(&:last).map do |k,v|
    [k.to_s.downcase.to_sym, v]
  end.to_h)

  scope :active, -> { where.not(status: :terminated) }

  def validate_number_in_range
    unless 1 <= number && number <= 5
      errors.add :number, 'インスタンス番号(number)が範囲外です。1以上5以下で指定してください。'
    end
  end

  def to_pb(team: false)
    Isuxportal::Proto::Resources::ContestantInstance.new(
      id: id,
      cloud_id: cloud_id,
      team_id: team_id,
      status: status_before_type_cast,
      number: number,
      public_ipv4_address: public_ipv4_address,
      private_ipv4_address: private_ipv4_address,
      team: team ? self.team&.to_pb : nil,
    )
  end
end
